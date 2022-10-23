import { RoomTypesInterface } from "./IRoomType";
import { EquipmentsInterface } from "./IEquipment";

export interface ResearchRoomsInterface {
    ID?: number,
    //Name?: string,

    RoomTypeID?: number,
    RoomType?: RoomTypesInterface,

    EquipmentID?: number,
    Equipment?: EquipmentsInterface,
    
}