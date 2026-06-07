import { QueryInterface} from "sequelize"
module.exports = {
  async up (queryInterface:QueryInterface) {
    await queryInterface.bulkInsert('hotels', [
      {
        name: 'Grand Palace Hotel',
        address: '123 MG Road, Bengaluru',
        location: 'Bengaluru, Karnataka',
        rating: 4.5,
        rating_count: 1250,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },
      {
        name: 'Ocean View Resort',
        address: '45 Beach Road, Goa',
        location: 'Goa',
        rating: 4.8,
        rating_count: 890,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },
      {
        name: 'Royal Heritage Inn',
        address: '12 Palace Street, Jaipur',
        location: 'Jaipur, Rajasthan',
        rating: 4.3,
        rating_count: 560,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },
      {
        name: 'Mountain Peak Lodge',
        address: '78 Hilltop Road, Manali',
        location: 'Manali, Himachal Pradesh',
        rating: 4.7,
        rating_count: 720,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },
      {
        name: 'Business Hub Suites',
        address: '99 Financial District, Hyderabad',
        location: 'Hyderabad, Telangana',
        rating: 4.1,
        rating_count: 340,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },
      {
        name: 'Lakeview Residency',
        address: '22 Lake Road, Udaipur',
        location: 'Udaipur, Rajasthan',
        rating: 4.6,
        rating_count: 675,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },
      {
        name: 'Airport Comfort Stay',
        address: '5 Airport Road, Chennai',
        location: 'Chennai, Tamil Nadu',
        rating: 3.9,
        rating_count: 210,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },
      {
        name: 'Green Valley Retreat',
        address: '15 Forest Lane, Coorg',
        location: 'Coorg, Karnataka',
        rating: 4.9,
        rating_count: 430,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },
      {
        name: 'City Center Hotel',
        address: '8 Connaught Place, New Delhi',
        location: 'New Delhi',
        rating: 4.2,
        rating_count: 980,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      },
      {
        name: 'Riverside Premium',
        address: '101 River Road, Kolkata',
        location: 'Kolkata, West Bengal',
        rating: 4.4,
        rating_count: 510,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null
      }
    ]);
  },

  async down (queryInterface:QueryInterface) {
    await queryInterface.bulkDelete('hotels', {});
  }
};
