export type GetAvailableRoomsDto = {
    roomCategoryId : number;
    checkInDate : string;
    checkOutDate : string;
}

export type updateBookingIdToRoomsDto = {
    bookingId : number;
    roomIds: number[];
}