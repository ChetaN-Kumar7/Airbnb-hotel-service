// step 1 - i want  to create from for certain set of dates and want 

import { CreationAttributes } from "sequelize";
import RoomCategory from "../db/models/roomCategory";
import { RoomGenerationJob } from "../dto/roomGeneration.dto";
import { RoomRepository } from "../repositories/room.repository";
import { RoomCategoryRepository } from "../repositories/roomCategory.repository";
import { BadRequestError, NotFoundError } from "../utils/errors/app.error";
import Room from "../db/models/room";
import logger from "../config/logger.config";

// creating object of room category class

const roomCategoryRepository = new RoomCategoryRepository();
const roomRepository = new RoomRepository()

export async function generateRooms(jobData:RoomGenerationJob) {

    let totalRoomsCreated =0
    let totalDatesProcessed =0 
    
    // now i have to check whether particular room category is valid or not in this case i need to check in db and for that i need repository for room category and in that i need findbyid function 

    const roomCategory = await roomCategoryRepository.findById(jobData.roomCategoryId)

    if(!roomCategory){
        throw new NotFoundError(`Room category with id ${jobData.roomCategoryId} not found`);
    }

    // now we have valid room Category now we check whether start date is greater then end date and start date is not current date it must be inn future

    const startDate = new Date(jobData.startDate)
    const endDate = new Date(jobData.endDate)

    if(startDate >= endDate){
        throw new BadRequestError(`Start date must be before end date`);
    }

    if(startDate< new Date()){
        throw new BadRequestError(`Start date must be in the future`);
    }

    const totalDays = Math.ceil((endDate.getTime()  - startDate.getTime())/(1000*60*60*24));

    logger.info(`Generating rooms for ${totalDays} days`);

    //now i have to process all the dates in batches 

    const batchSize = jobData.batchSize || 100

    //now i have to traverse over the batch size to cover all the dates

    const currentDate = new Date(startDate)

    while(currentDate < endDate){
        const batchEndDate = new Date(currentDate)

        batchEndDate.setDate(batchEndDate.getDate()+batchSize);

        if(batchEndDate > endDate){
            batchEndDate.setTime(endDate.getTime())
        }

        // now i have current date and batch end date now we have to process these batches

        const batchResult = await processDateBatch(roomCategory,currentDate,batchEndDate,jobData.priceOverride)

        totalRoomsCreated += batchResult.roomsCreated;
        totalDatesProcessed += batchResult.datesProcessed;

        currentDate.setTime(batchEndDate.getTime())
    }

    return{
        totalDatesProcessed,
        totalRoomsCreated,
    }
}

export async function processDateBatch(roomCategory:RoomCategory, startDate: Date,endDate:Date,priceOverride?: number) {
    // first i have to check whether for particular roomCategory id and for those dates room already exist or not and for that i have to query the db and for that  we need repository layer for room

    // now i have to check for each date from start to end whether room is existed or not for that i have to do looping 

    let roomsCreated = 0;
    let datesProcessed = 0;

    const roomsToCreate:CreationAttributes<Room>[] = []; 

    let currentDate = new Date(startDate)

    while(currentDate <= endDate){
        
        const existingRoom = await roomRepository.findByRoomCategoryIdAndDate(roomCategory.id,currentDate)

        if(!existingRoom){
            const roomPayload = {
                hotelId: roomCategory.hotelId,
                roomCategoryId: roomCategory.id,
                dateOfAvailability: new Date(currentDate),
                price: priceOverride || roomCategory.price,
                createdAt: new Date(),
                updatedAt: new Date(),
                deletedAt: null,
            }

            roomsToCreate.push(roomPayload)
        }

        currentDate.setDate(currentDate.getDate() + 1);
        datesProcessed++;
    }

    // now we have ready with our roomsToCreate array now we check the length of array if its greater then 0 then we make a db call for bulk insert

    if(roomsToCreate.length>0){
        await roomRepository.bulkCreate(roomsToCreate)
        roomsCreated += roomsToCreate.length
    }

    return {
        roomsCreated,datesProcessed
    }

}