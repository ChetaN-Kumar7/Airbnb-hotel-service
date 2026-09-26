import { serverConfig } from "../config";

export const getAvailableRooms = async (
  roomCategoryId: number,
  checkInDate: string,
  checkOutDate: string
) => {
  const params = new URLSearchParams({
    roomCategoryId: String(roomCategoryId),
    checkInDate,
    checkOutDate,
  });

  console.log("Room availability request:", {
    roomCategoryId,
    checkInDate,
    checkOutDate,
  });

  console.log(
    "Request URL:",
    `${serverConfig.HOTEL_SERVICE_URL}rooms/available?${params}`
  );

  const response = await fetch(
    `${serverConfig.HOTEL_SERVICE_URL}rooms/available?${params}`
  );

  if (!response.ok) {
    throw new Error(`Request failed: ${response.status}`);
  }

  return response.json();
};

export const updateBookingIdToRooms = async (
  bookingId: number,
  roomIds: number[],
) => {
    console.log(
    "Request URL:",
    `${serverConfig.HOTEL_SERVICE_URL}rooms/update-booking-id`
  );
    const response = await fetch(`${serverConfig.HOTEL_SERVICE_URL}rooms/update-booking-id`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            bookingId,
            roomIds,
        }),
    });
    if (!response.ok) {
        throw new Error(`Request failed: ${response.status}`);
    }

    return response.json();
}