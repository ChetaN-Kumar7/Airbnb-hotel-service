import express from 'express';
import {  validateRequestBody } from '../../validators';
import { hotelSchema } from '../../validators/hotel.validator';
import { createHotelHandler, deleteHotelHandler, getAllHotelHandler, getHotelByIdcHandler } from '../../controllers/hotel.controller';


const hotelRouter = express.Router();

hotelRouter.post('/', validateRequestBody(hotelSchema), createHotelHandler); // TODO: Resolve this TS compilation issue

hotelRouter.get('/:id', getHotelByIdcHandler);

hotelRouter.get('/', getAllHotelHandler);
// hotelRouter.patch('/:id',updateHotelHandler);
hotelRouter.delete('/:id',deleteHotelHandler);

export default hotelRouter;