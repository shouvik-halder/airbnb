import {
  Model,
  InferAttributes,
  InferCreationAttributes,
  CreationOptional,
  DataTypes
} from "sequelize";
import { sequelize } from "./sequelize";

class RoomType extends Model<
  InferAttributes<RoomType>,
  InferCreationAttributes<RoomType>
> {
  declare id: CreationOptional<number>;
  declare hotel_id: number;

  declare name: string;
  declare description: CreationOptional<string | null>;
  declare max_occupancy: number;
  declare room_count: number;

  declare createdAt: CreationOptional<Date>;
  declare updatedAt: CreationOptional<Date>;
  declare deletedAt: CreationOptional<Date | null>;
}

RoomType.init(
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
    name: {
      type: DataTypes.STRING(100),
      allowNull: false,
    },
    description: {
      type: DataTypes.TEXT,
      allowNull: true,
    },
    max_occupancy: {
      type: DataTypes.INTEGER,
      allowNull: false,
      validate: {
        min: 1,
      },
    },
    room_count:{
      type: DataTypes.INTEGER,
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
    tableName: "room_types",
    timestamps: true,
    paranoid: true,
    underscored: true,
  }
);

export default RoomType;