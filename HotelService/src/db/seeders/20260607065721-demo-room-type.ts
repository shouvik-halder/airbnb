import { QueryInterface} from "sequelize"
module.exports = {
  async up (queryInterface:QueryInterface) {
   await queryInterface.bulkInsert('room_types', [
      // Hotel 1
      {
        hotel_id: 1,
        name: 'Standard Room',
        description: 'Comfortable room with queen bed and city view.',
        max_occupancy: 2,
        room_count: 100,
        base_price: 2999.00,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },
      {
        hotel_id: 1,
        name: 'Deluxe Room',
        description: 'Spacious room with premium amenities.',
        max_occupancy: 3,
        room_count: 50,
        base_price: 4999.00,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },
      {
        hotel_id: 1,
        name: 'Suite',
        description: 'Luxury suite with separate living area.',
        max_occupancy: 4,
        room_count: 20,
        base_price: 8999.00,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },

      // Hotel 2
      {
        hotel_id: 2,
        name: 'Standard Room',
        description: 'Beach-facing standard room.',
        max_occupancy: 2,
        room_count: 80,
        base_price: 3499.00,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },
      {
        hotel_id: 2,
        name: 'Ocean View Deluxe',
        description: 'Premium room with ocean view.',
        max_occupancy: 3,
        room_count: 40,
        base_price: 6499.00,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },

      // Hotel 3
      {
        hotel_id: 3,
        name: 'Heritage Room',
        description: 'Traditional Rajasthani themed room.',
        max_occupancy: 2,
        room_count: 60,
        base_price: 3999.00,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },
      {
        hotel_id: 3,
        name: 'Royal Suite',
        description: 'Luxury suite with royal interiors.',
        max_occupancy: 5,
        room_count: 15,
        base_price: 12999.00,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },

      // Hotel 4
      {
        hotel_id: 4,
        name: 'Mountain View Room',
        description: 'Room overlooking the mountains.',
        max_occupancy: 2,
        room_count: 70,
        base_price: 4499.00,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },
      {
        hotel_id: 4,
        name: 'Family Suite',
        description: 'Ideal for families and groups.',
        max_occupancy: 6,
        room_count: 25,
        base_price: 9999.00,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },

      // Hotel 5
      {
        hotel_id: 5,
        name: 'Business Room',
        description: 'Designed for business travellers.',
        max_occupancy: 2,
        room_count: 90,
        base_price: 3799.00,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },
      {
        hotel_id: 5,
        name: 'Executive Suite',
        description: 'Premium suite with work area.',
        max_occupancy: 4,
        room_count: 30,
        base_price: 8499.00,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      }
    ]);
  },

  async down (queryInterface:QueryInterface) {
    
     await queryInterface.bulkDelete('room_types', {});
     
  }
};
