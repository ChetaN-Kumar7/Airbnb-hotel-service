import  PrismaClient from "../prisma/client";
import { CreateBookingDto } from "../dto/booking.dto";
import { confirmBooking, createBooking, createIdempotencyKey, finalizeIdempotencyKey, getIdempotencyKeyWithLock } from "../repositries/booking.repository";
import { BadRequestError, InternalServerError, NotFoundError } from "../utils/errors/app.error";
import { generateIdempotencyKey } from "../utils/generateIdempotencyKey";
import { serverConfig } from "../config";
import { redlock } from "../config/redis.config";

export  async function createBookingService(createBookingDto: CreateBookingDto) {

    const ttl = serverConfig.LOCK_TTL;
    const bookingResource = `hotel:${createBookingDto.hotelId}`

    try{
        await redlock.acquire([bookingResource], ttl);
            const booking = await createBooking({
            userId:createBookingDto.userId,
            hotelId:createBookingDto.hotelId,
            bookingAmount : createBookingDto.bookingAmount,
            totalGuests : createBookingDto.totalGuests
        })

        const idempotencyKey = generateIdempotencyKey();

        await createIdempotencyKey(idempotencyKey,booking.id);

        return {
            bookingId: booking.id,
            idempotencyKey : idempotencyKey,
        }
    }
    catch(error){
        throw new InternalServerError('Failed to acquire lock for resource')
    }

    
}

export async function confirmBookingService(idempotencyKey: string) {

    return await PrismaClient.$transaction( async(tx)=>{
        const idempotencyKeyData = await getIdempotencyKeyWithLock(tx,idempotencyKey)

        if(!idempotencyKeyData || !idempotencyKeyData.bookingId){
            throw new NotFoundError('Idempotency key not found');
        }

        if(idempotencyKeyData.finalized){
            throw new BadRequestError('Idempotency key already finalized');
        }

        const  booking = await confirmBooking(tx,idempotencyKeyData.bookingId);
        await finalizeIdempotencyKey(tx,idempotencyKey)

        return booking;
    })

}