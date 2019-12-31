-- phpMyAdmin SQL Dump
-- version 4.8.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 31 Des 2019 pada 11.07
-- Versi server: 10.1.37-MariaDB
-- Versi PHP: 7.3.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `zainab_api`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `zainab_author`
--

CREATE TABLE `zainab_author` (
  `id` int(100) NOT NULL,
  `nama` varchar(100) NOT NULL,
  `alamat` varchar(500) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(50) NOT NULL,
  `email` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `zainab_author`
--

INSERT INTO `zainab_author` (`id`, `nama`, `alamat`, `username`, `password`, `email`) VALUES
(1, 'Adam J. Whitworth', '461 Dogwood Road', 'Hishowas', 'uujeed3Ae', 'AdamJWhitworth@teleworm.us'),
(2, 'Susan C. Groth', '3885 Hershell Hollow Road', 'Tobjew', 'Jain7Ohghiu', 'SusanCGroth@teleworm.us'),
(3, 'Stanley T. Stevenson', '1418 Red Dog Road', 'Brold1961', 'ooGeitohj9', 'StanleyTStevenson@armyspy.com');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `zainab_author`
--
ALTER TABLE `zainab_author`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `zainab_author`
--
ALTER TABLE `zainab_author`
  MODIFY `id` int(100) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
