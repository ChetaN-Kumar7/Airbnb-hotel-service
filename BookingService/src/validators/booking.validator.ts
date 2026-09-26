import {z} from 'zod';

export const createBookingSchema = z.object({
    userId : z.number({message:"User Id must be present"}),
    hotelId: z.number({message:"Hotel Id must be present"}),
    totalGuests : z.number({message:"Total guest must be present"}).min(1,{message:"Total guest must be at least 1"}),
    bookingAmount:  z.number({message:"Booking amount must be present"}).min(1,{message:"Booking amount must be greater then 1"}),
    checkInDate: z.string({message: "check-in date must be present"}),
    checkOutDate: z.string({message: "check-in date must be present"}),
    roomCategoryId: z.number({message: "Room category id must be present"}),
})