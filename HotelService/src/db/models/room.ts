import {
  Model,
  InferAttributes,
  InferCreationAttributes,
  CreationOptional,
  DataTypes
} from "sequelize";
import { sequelize } from "./sequelize";

export class Room extends Model<
  InferAttributes<Room>,
  InferCreationAttributes<Room>
> {
  declare id: CreationOptional<number>;
  declare hotel_id: number;
  declare room_type_id: number;   // or number if using room_type_id
  declare room_number: string;
  declare floor:number;
  declare status:RoomStatus;
  declare statusDate: Date;

  declare createdAt: CreationOptional<Date>;
  declare updatedAt: CreationOptional<Date>;
  declare deletedAt: CreationOptional<Date | null>;
}

export enum RoomStatus{
    available = "available",
    occupied = "occupied",
    maintenance = "maintenance"
}

Room.init(
  {
    id: {
      type: DataTypes.INTEGER,
      autoIncrement: true,
      primaryKey: true,
    },
    hotel_id: {
      type: DataTypes.INTEGER,
      allowNull: false,
    },
    room_type_id: {
      type: DataTypes.INTEGER,
      allowNull: false,
    },
    room_number: {
      type: DataTypes.STRING(20),
      allowNull: false,
    },
    floor: {
      type: DataTypes.INTEGER,
      allowNull: true,
    },
    status: {
      type: DataTypes.ENUM,
      allowNull: false,
      defaultValue: RoomStatus.available,
      values: [...Object.values(RoomStatus)]
    },
    statusDate:{
      type: "DATE",
      allowNull: false
    },
    createdAt:{
        type:"DATE",
        defaultValue:new Date()
    },
    updatedAt:{
        type:"DATE",
        defaultValue:new Date()
    },
    deletedAt:{
        type:"DATE",
        defaultValue:null
    }
  },
  {
    sequelize,
    tableName: "rooms",
    timestamps: true,
    underscored: true, // maps created_at instead of createdAt
  }
);

export default Room;