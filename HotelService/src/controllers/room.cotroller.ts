import { NextFunction, Request, Response } from "express";
import { getAvailableRoomService, updateBookingIdToRoomsService } from "../services/room.service";


export async function getAvailableRoomHandler(req: Request, res:Response , next:NextFunction) {
    
    console.log("Yha aya");
    const rooms = await getAvailableRoomService({roomCategoryId:Number(req.query.roomCategoryId),checkInDate:req.query.checkInDate as string,checkOutDate: req.query.checkOutDate as string});
    res.status(200).json({
        message: "Rooms found successfully",
        data: rooms,
        success: true,
    })
    
}


export async function updateBookingIdToRoomsHandler(req: Request, res:Response , next:NextFunction) {

    const rooms = await updateBookingIdToRoomsService(req.body);
    res.status(200).json({
        message: "Booking Id updated to rooms successfully",
        data: rooms,
        success: true,
    })
    
}
