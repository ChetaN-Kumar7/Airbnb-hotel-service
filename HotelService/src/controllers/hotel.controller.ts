import { NextFunction, Request, Response } from "express";
import { createHotelService, getAllHotelsService, getHotelByIdService, updateHotelService } from "../services/hotel.service";


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

export async function updateHotelHandler(req:Request, res:Response, next:NextFunction) {
    const hotelResponse = await updateHotelService(Number(req.params.id),req.body);
    res.status(201).json({
        message: "Hotels founded successfully",
        data: hotelResponse,
        success: true,
    })
}
