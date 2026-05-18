import logger from "../config/logger.config";
import Hotel from "../db/models/hotel";
import { createHotelDTO, updateDTO } from "../dto/hotel.dto";
import { NotFoundError } from "../utils/errors/app.error";

export async function createHotel(hotelData:createHotelDTO){
    const hotel = await Hotel.create({
        name:hotelData.name,
        address:hotelData.address,
        location:hotelData.location,
        rating: hotelData.rating,
        ratingCount:hotelData.ratingCount,
    });
    logger.info(`hotel created successfully ${hotel.id}`)
    return hotel;
}

export async function getHotelById(id:number) {
    const hotel = await Hotel.findByPk(id)

    if(!hotel){
        logger.error(`hotel not found ${id}`)
        throw new NotFoundError(`hotel with ${id} not found`)
    }

    logger.info(`hotel found:${hotel.id}`)
    return hotel;
}

export async function getAllHotels() {
    const hotels = await Hotel.findAll()

    if(!hotels){
        logger.error(`hotel not found`)
        throw new NotFoundError(`hotels not found`)
    }

    logger.info(`hotels founded ${hotels}`)
    return hotels;
}

export async function updateHotel(id: number,hotelData: updateDTO) {
   
  const hotel = await Hotel.findByPk(id);

  if (!hotel) {
    throw new NotFoundError(
      `Hotel with id ${id} not found`
    );
  }

  return await hotel.update(hotelData);
}
