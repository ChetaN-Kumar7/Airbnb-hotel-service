import { GetAvailableRoomsDto, updateBookingIdToRoomsDto } from "../dto/room.dto";
import { RoomRepository } from "../repositories/room.repository";

const roomRepository = new RoomRepository();

export async function getAvailableRoomService(getAvailableRoomsDto : GetAvailableRoomsDto) {
    const rooms = await roomRepository.findByRoomCategoryIdAndDateRange(getAvailableRoomsDto.roomCategoryId,new Date(getAvailableRoomsDto.checkInDate),new Date(getAvailableRoomsDto.checkOutDate));
    return rooms;
}

export async function updateBookingIdToRoomsService(updateBookingIdToRoomsDto:updateBookingIdToRoomsDto) {
    return await roomRepository.updateBookingIdToRooms(updateBookingIdToRoomsDto.bookingId,updateBookingIdToRoomsDto.roomIds)
}