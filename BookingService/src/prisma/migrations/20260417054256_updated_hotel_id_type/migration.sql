/*
  Warnings:

  - You are about to alter the column `hotelId` on the `Booking` table. The data in that column could be lost. The data in that column will be cast from `VarChar(191)` to `Int`.

*/
-- AlterTable
ALTER TABLE `Booking` MODIFY `hotelId` INTEGER NOT NULL;
