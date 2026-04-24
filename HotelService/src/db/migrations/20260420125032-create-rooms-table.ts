import { QueryInterface } from "sequelize";
module.exports = {
  async up (queryInterface:QueryInterface) {
    await queryInterface.sequelize.query(`
      CREATE TABLE IF NOT EXISTS rooms (
    id              INT AUTO_INCREMENT PRIMARY KEY,
    room_type_id    INT NOT NULL,

    room_number     VARCHAR(20) NOT NULL,
    floor           INT,
    status          ENUM('available', 'occupied', 'maintenance') NOT NULL DEFAULT 'available',
    status_date     DATE NOT NULL,
    

    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP NULL DEFAULT NULL,

    CONSTRAINT fk_rooms_room_type
        FOREIGN KEY (room_type_id)
        REFERENCES room_types(id)
        ON DELETE CASCADE,

    UNIQUE KEY unique_room_number (room_type_id, room_number)
)`
)
  },

  async down (queryInterface:QueryInterface) {
    await queryInterface.sequelize.query(`
      DROP TABLE IF EXISTS rooms`)
  }
};
