import { get } from './http';

export const carQuery = (data) => get(`/api/car`, data);