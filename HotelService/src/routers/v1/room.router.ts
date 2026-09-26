import express from 'express';
import {  validateQueryParams, validateRequestBody } from '../../validators';
import { getAvailableRoomHandler, updateBookingIdToRoomsHandler,  } from '../../controllers/room.cotroller';
import { getAvailableRoomsSchema, updateBookingIdToRoomsSchema } from '../../validators/room.validator';

const roomRouter = express.Router();

roomRouter.get('/available',validateQueryParams(getAvailableRoomsSchema), getAvailableRoomHandler );
roomRouter.post('/update-booking-id',validateRequestBody(updateBookingIdToRoomsSchema), updateBookingIdToRoomsHandler );


export default roomRouter;