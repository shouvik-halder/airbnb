import {QueryInterface} from 'sequelize';
module.exports = {
  async up (queryInterface:QueryInterface) {
    await queryInterface.sequelize.query(`
      ALTER TABLE rooms
      ADD COLUMN booking_id INT NULL;`)
  },

  async down (queryInterface:QueryInterface) {
    await queryInterface.sequelize.query(`
      ALTER TABLE rooms
      DROP COLUMN booking_id;`)
  }
};
