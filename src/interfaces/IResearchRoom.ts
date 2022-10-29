import { RoomTypesInterface } from "./IRoomType";

export interface ResearchRoomsInterface {
    ID?: number,
    Name: string,

    RoomTypeID?: number,
    RoomType?: RoomTypesInterface,
}