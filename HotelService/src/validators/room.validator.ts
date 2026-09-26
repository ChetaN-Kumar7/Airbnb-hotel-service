import z from "zod";


export const getAvailableRoomsSchema = z.object({
    roomCategoryId: z.coerce.number({message: "Room category id must be present"}),
    checkInDate : z.string({message:"Check-in date must be present"}),
    checkOutDate : z.string({message:"Check-out date must be present"}),
})

export const updateBookingIdToRoomsSchema = z.object({
    bookingId: z.number({message: "Booking id must be present"}),
    roomIds : z.array(z.number({message:"Rooms id must be present"})),
})