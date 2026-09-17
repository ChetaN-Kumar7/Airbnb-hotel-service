import { CreationAttributes } from "sequelize";
import Room from "../db/models/room";
import BaseRepository from "./base.repository";

export class RoomRepository extends BaseRepository<Room>{
    constructor() {
        super(Room);
    }

    async findByRoomCategoryIdAndDate(roomCategoryId: number,currentDate: Date) {

        return await Room.findOne({
            where : {
                roomCategoryId: roomCategoryId,
                dateOfAvailability: currentDate,
                deletedAt: null
            }
        })
        
    }

    async bulkCreate(rooms : CreationAttributes<Room>[]){
        return Room.bulkCreate(rooms)
    }
}
