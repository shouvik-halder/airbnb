import { QueryInterface } from "sequelize";

module.exports = {
  async up(queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
      CREATE TABLE IF NOT EXISTS room_types (
    id              INT AUTO_INCREMENT PRIMARY KEY,
    hotel_id        INT NOT NULL,

    name            VARCHAR(100) NOT NULL,
    description     TEXT,
    max_occupancy   INT NOT NULL CHECK (max_occupancy > 0),
    room_count      INT,

    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT fk_room_types_hotel
        FOREIGN KEY (hotel_id)
        REFERENCES hotels(id)
        ON DELETE CASCADE,

    UNIQUE KEY unique_hotel_roomtype (hotel_id, name)
) 
      `);
  },

  async down(queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
      DROP TABLE IF NOT EXISTS room_types`)
  },
};
