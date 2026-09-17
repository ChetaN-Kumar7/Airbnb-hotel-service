import logger from "../config/logger.config";
import Hotel from "../db/models/hotel";
import { NotFoundError } from "../utils/errors/app.error";
import BaseRepository from "./base.repository";

export class HotelRepository extends BaseRepository<Hotel> {
    constructor() {
        super(Hotel);
    }

    async findAll() {
        const hotels = await this.model.findAll({
            where: {
                deleted_at: null
            }
        });

        if (!hotels) {
            logger.error(`No hotels found`);
            throw new NotFoundError(`No hotels found`);
        }

        logger.info(`Hotels found: ${hotels.length}`);
        return hotels;
    }

    async softDelete(id: number) {
        const hotel = await Hotel.findByPk(id);

        if(!hotel) {
            logger.error(`Hotel not found: ${id}`);
            throw new NotFoundError(`Hotel with id ${id} not found`);
        }

        hotel.deleted_at= new Date();
        await hotel.save(); // Save the changes to the database
        logger.info(`Hotel soft deleted: ${hotel.id}`);
        return true;
    }

}