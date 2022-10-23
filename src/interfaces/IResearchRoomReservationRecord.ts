import { ResearchRoomsInterface } from "./IResearchRoom";
import { UsersInterface } from "./IUser";
import { AddOnsInterface } from "./IAddOn";
import { TimeRoomsInterface } from "./ITimeRoom";
//belong ResearchRooms
import { EquipmentsInterface } from "./IEquipment";
import { RoomTypesInterface } from "./IRoomType";

export interface ResearchRoomReservationRecordInterface {
    ID?: number,
    BookDate: Date | null,
    ResearchRoomID?: number,
    ResearchRoom?: ResearchRoomsInterface,
   
    UserID?: number,
    User?: UsersInterface,

    AddOnID?: number,
    AddOn?: AddOnsInterface,

    TimeRoomID?: number,
    TimeRoom?: TimeRoomsInterface,

    RoomTypeID?: number,
    RoomType?: RoomTypesInterface,

    EquipmentID?: number,
    Equipment?: EquipmentsInterface
    
  }