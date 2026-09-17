import { NextFunction, Request, Response } from "express";
import { createHotelService, deleteHotelService, getAllHotelsService, getHotelByIdService } from "../services/hotel.service";



export async function createHotelHandler(req:Request, res:Response, next:NextFunction) {
    const hotelResponse = await createHotelService(req.body);
    res.status(201).json({
        message: "Hotel created successfully",
        data: hotelResponse,
        success: true,
    })
}

export async function getHotelByIdcHandler(req:Request, res:Response, next:NextFunction) {
    const hotelResponse = await getHotelByIdService(Number(req.params.id));
    res.status(201).json({
        message: "Hotel founded successfully",
        data: hotelResponse,
        success: true,
    })
}

export async function getAllHotelHandler(req:Request, res:Response, next:NextFunction) {
    const hotelResponse = await getAllHotelsService();
    res.status(201).json({
        message: "Hotels founded successfully",
        data: hotelResponse,
        success: true,
    })
}

// export async function updateHotelHandler(req:Request, res:Response, next:NextFunction) {
//     const hotelResponse = await updateHotelService(Number(req.params.id),req.body);
//     res.status(201).json({
//         message: "Hotels founded successfully",
//         data: hotelResponse,
//         success: true,
//     })
// }

export async function deleteHotelHandler(req:Request,res:Response,next:NextFunction) {
    const hotelResponse = await deleteHotelService(Number(req.params.id));
    res.status(201).json({
        message: "Hotels deleted successfully",
        data: hotelResponse,
        success: true,
    })
}
