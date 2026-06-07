import {QueryInterface} from 'sequelize'
module.exports = {
  async up (queryInterface:QueryInterface) {
    await queryInterface.sequelize.query(
      ` ALTER TABLE room_types
        ADD COLUMN base_price DECIMAL(10, 2) NOT NULL DEFAULT 0.00 AFTER room_count;`
    );
  },

  async down (queryInterface:QueryInterface) {
    queryInterface.sequelize.query(
      `ALTER TABLE room_types
       DROP COLUMN base_price;`
    )
  }
};
