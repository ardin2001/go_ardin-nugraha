-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 19 Mar 2023 pada 04.23
-- Versi server: 10.4.27-MariaDB
-- Versi PHP: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `alta_online_shop`
--

DELIMITER $$
--
-- Fungsi
--
CREATE DEFINER=`root`@`localhost` FUNCTION `deleteTransaction` (`t_id` INT) RETURNS INT(11) DETERMINISTIC BEGIN
DELETE FROM transactions WHERE id=t_id;
RETURN 1;
END$$

CREATE DEFINER=`root`@`localhost` FUNCTION `deleteTransdetail` (`td_id` INT) RETURNS INT(11) DETERMINISTIC BEGIN
DELETE FROM transaction_details WHERE id=td_id;
RETURN 1;
END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Struktur dari tabel `operators`
--

CREATE TABLE `operators` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `operators`
--

INSERT INTO `operators` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'pt maju jaya', '2023-03-17 04:16:11', '2023-03-17 04:16:11'),
(2, 'pt maju mundur', '2023-03-17 04:16:24', '2023-03-17 04:16:24'),
(3, 'pt alfaria jaya', '2023-03-17 04:16:41', '2023-03-17 04:16:41'),
(4, 'pt lancar jaya', '2023-03-17 04:16:55', '2023-03-17 04:16:55');

-- --------------------------------------------------------

--
-- Struktur dari tabel `payment_methods`
--

CREATE TABLE `payment_methods` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `status` smallint(11) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `payment_methods`
--

INSERT INTO `payment_methods` (`id`, `name`, `status`, `created_at`, `updated_at`) VALUES
(1, 'OVO', 1, '2023-03-17 04:44:39', '2023-03-17 04:44:39'),
(2, 'gopay', 1, '2023-03-17 04:44:44', '2023-03-17 04:44:44'),
(3, 'ATM', 1, '2023-03-17 04:44:48', '2023-03-17 04:44:48'),
(4, 'dana', 1, '2023-03-17 04:44:53', '2023-03-17 04:44:53');

-- --------------------------------------------------------

--
-- Struktur dari tabel `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `product_type_id` int(11) NOT NULL,
  `operator_id` int(11) NOT NULL,
  `code` varchar(50) NOT NULL,
  `name` varchar(100) NOT NULL,
  `status` smallint(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `products`
--

INSERT INTO `products` (`id`, `product_type_id`, `operator_id`, `code`, `name`, `status`, `created_at`, `updated_at`) VALUES
(1, 1, 3, 'dcbhgbw732bdyu', 'product dummy', 1, '2023-03-17 04:22:41', '2023-03-17 04:22:41'),
(2, 1, 3, 'hr7hgbw732bdyu', 'laptop', 1, '2023-03-17 04:23:00', '2023-03-17 04:23:00'),
(3, 2, 1, 'gjghftw732bdyu', 'wajan', 1, '2023-03-17 04:30:11', '2023-03-17 04:30:11'),
(4, 2, 1, 'otunfftw732bdyu', 'piring', 1, '2023-03-17 04:30:32', '2023-03-17 04:30:32'),
(5, 2, 1, 'urt4fftw732bdyu', 'panci', 1, '2023-03-17 04:30:49', '2023-03-17 04:30:49'),
(6, 3, 4, 'ir64rfftw732bdyu', 'baju evos', 1, '2023-03-17 04:31:01', '2023-03-17 04:31:01'),
(7, 3, 4, 'jdf8erfftw732bdyu', 'baju rrq', 1, '2023-03-17 04:31:12', '2023-03-17 04:31:12'),
(8, 3, 4, 'nt8derfftw732bdyu', 'baju alter ego', 1, '2023-03-17 04:31:23', '2023-03-17 04:31:23');

-- --------------------------------------------------------

--
-- Struktur dari tabel `product_descriptions`
--

CREATE TABLE `product_descriptions` (
  `id` int(11) NOT NULL,
  `description` text NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `product_descriptions`
--

INSERT INTO `product_descriptions` (`id`, `description`, `created_at`, `updated_at`) VALUES
(1, 'ini adalah barang elektronik berkualitas bagus', '2023-03-17 04:39:16', '2023-03-17 04:39:16'),
(2, 'ini adalah barang laptop berkualitas tinggi', '2023-03-17 04:39:23', '2023-03-17 04:39:23'),
(3, 'ini adalah wajan dengan bahan bagus', '2023-03-17 04:39:30', '2023-03-17 04:39:30'),
(4, 'ini adalah piring dengan bahan terbaik', '2023-03-17 04:39:44', '2023-03-17 04:39:44'),
(5, 'ini adalah panji dengan bahan premium', '2023-03-17 04:39:52', '2023-03-17 04:39:52'),
(6, 'ini adalah baju dengan bahan halus', '2023-03-17 04:40:00', '2023-03-17 04:40:00'),
(7, 'ini adalah baju dengan bahan premium', '2023-03-17 04:40:10', '2023-03-17 04:40:10'),
(8, 'ini adalah baju dengan kualitas ori', '2023-03-17 04:40:20', '2023-03-17 04:40:20');

-- --------------------------------------------------------

--
-- Struktur dari tabel `product_types`
--

CREATE TABLE `product_types` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `product_types`
--

INSERT INTO `product_types` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'barang elektronik', '2023-03-17 04:14:26', '2023-03-17 04:14:26'),
(2, 'peralatan masak', '2023-03-17 04:14:45', '2023-03-17 04:14:45'),
(3, 'pakaian', '2023-03-17 04:14:56', '2023-03-17 04:14:56'),
(4, 'mainan', '2023-03-17 04:15:10', '2023-03-17 04:15:10'),
(5, 'kartu provider', '2023-03-17 04:15:21', '2023-03-17 04:15:21');

-- --------------------------------------------------------

--
-- Struktur dari tabel `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `payment_method_id` int(11) NOT NULL,
  `status` varchar(10) NOT NULL,
  `total_qty` int(11) NOT NULL,
  `total_price` int(252) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `transactions`
--

INSERT INTO `transactions` (`id`, `user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 'lunas', 8, 150000, '2023-03-17 05:12:34', '2023-03-17 05:12:34'),
(2, 1, 2, 'lunas', 5, 100000, '2023-03-17 05:12:34', '2023-03-17 05:12:34'),
(3, 1, 3, 'lunas', 5, 300000, '2023-03-17 05:12:34', '2023-03-17 05:12:34'),
(4, 2, 3, 'lunas', 5, 300000, '2023-03-17 05:12:34', '2023-03-17 05:12:34'),
(5, 2, 2, 'lunas', 6, 200000, '2023-03-17 05:12:34', '2023-03-17 05:12:34'),
(6, 2, 1, 'lunas', 6, 200000, '2023-03-17 05:12:34', '2023-03-17 05:12:34'),
(7, 3, 3, 'lunas', 6, 500000, '2023-03-17 05:12:34', '2023-03-17 05:12:34'),
(8, 3, 2, 'lunas', 7, 300000, '2023-03-17 05:12:34', '2023-03-17 05:12:34'),
(9, 3, 2, 'lunas', 4, 200000, '2023-03-17 05:12:35', '2023-03-17 05:12:35'),
(10, 4, 3, 'lunas', 4, 200000, '2023-03-17 05:12:35', '2023-03-17 05:12:35'),
(11, 4, 3, 'lunas', 4, 400000, '2023-03-17 05:12:35', '2023-03-17 05:12:35'),
(12, 4, 3, 'lunas', 5, 200000, '2023-03-17 05:12:35', '2023-03-17 05:12:35'),
(13, 5, 1, 'lunas', 5, 200000, '2023-03-17 05:12:35', '2023-03-17 05:12:35'),
(14, 5, 1, 'lunas', 5, 400000, '2023-03-17 05:12:35', '2023-03-17 05:12:35');

--
-- Trigger `transactions`
--
DELIMITER $$
CREATE TRIGGER `deleteTransactionTrigger` BEFORE DELETE ON `transactions` FOR EACH ROW BEGIN
DECLARE trans_id INT;
SET trans_id = OLD.id;
DELETE FROM transaction_details WHERE transaction_id = trans_id;
END
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaction_details`
--

CREATE TABLE `transaction_details` (
  `id` int(11) NOT NULL,
  `transaction_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `status` varchar(10) NOT NULL,
  `qty` int(11) NOT NULL,
  `price` int(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `transaction_details`
--

INSERT INTO `transaction_details` (`id`, `transaction_id`, `product_id`, `status`, `qty`, `price`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 'lunas', 3, 150000, '2023-03-17 05:26:08', '2023-03-17 05:26:08'),
(2, 1, 2, 'lunas', 2, 350000, '2023-03-17 05:26:08', '2023-03-17 05:26:08'),
(3, 1, 4, 'lunas', 3, 350000, '2023-03-17 05:26:08', '2023-03-17 05:26:08'),
(4, 2, 4, 'lunas', 2, 350000, '2023-03-17 05:26:08', '2023-03-17 05:26:08'),
(5, 2, 5, 'lunas', 2, 350000, '2023-03-17 05:26:08', '2023-03-17 05:26:08'),
(6, 2, 6, 'lunas', 4, 150000, '2023-03-17 05:26:08', '2023-03-17 05:26:08'),
(7, 3, 1, 'lunas', 3, 450000, '2023-03-17 05:26:08', '2023-03-17 05:26:08'),
(8, 3, 3, 'lunas', 1, 45000, '2023-03-17 05:26:08', '2023-03-17 05:26:08'),
(9, 3, 6, 'lunas', 1, 50000, '2023-03-17 05:26:08', '2023-03-17 05:26:08'),
(10, 4, 2, 'lunas', 2, 200000, '2023-03-17 05:26:08', '2023-03-17 05:26:08'),
(11, 4, 5, 'lunas', 2, 400000, '2023-03-17 05:26:08', '2023-03-17 05:26:08'),
(12, 4, 7, 'lunas', 1, 100000, '2023-03-17 05:26:08', '2023-03-17 05:26:08'),
(13, 5, 7, 'lunas', 2, 500000, '2023-03-17 05:26:08', '2023-03-17 05:26:08'),
(14, 5, 3, 'lunas', 1, 300000, '2023-03-17 05:26:08', '2023-03-17 05:26:08'),
(15, 5, 1, 'lunas', 3, 500000, '2023-03-17 05:26:08', '2023-03-17 05:26:08'),
(16, 6, 3, 'lunas', 2, 50000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(17, 6, 1, 'lunas', 3, 50000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(18, 6, 4, 'lunas', 1, 50000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(19, 7, 6, 'lunas', 1, 100000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(20, 7, 7, 'lunas', 2, 100000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(21, 7, 3, 'lunas', 1, 100000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(22, 8, 1, 'lunas', 3, 100000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(23, 8, 3, 'lunas', 1, 50000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(24, 8, 5, 'lunas', 3, 50000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(25, 9, 2, 'lunas', 2, 100000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(26, 9, 4, 'lunas', 1, 200000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(27, 9, 8, 'lunas', 1, 400000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(28, 10, 4, 'lunas', 1, 50000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(29, 10, 3, 'lunas', 2, 100000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(30, 10, 5, 'lunas', 2, 200000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(31, 11, 4, 'lunas', 1, 200000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(32, 11, 3, 'lunas', 2, 50000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(33, 11, 5, 'lunas', 2, 100000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(34, 12, 6, 'lunas', 2, 400000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(35, 12, 3, 'lunas', 2, 50000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(36, 12, 5, 'lunas', 1, 40000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(37, 13, 3, 'lunas', 1, 50000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(38, 13, 5, 'lunas', 2, 100000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(39, 13, 4, 'lunas', 2, 50000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(40, 14, 1, 'lunas', 3, 50000, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(41, 14, 3, 'lunas', 2, 25000, '0000-00-00 00:00:00', '0000-00-00 00:00:00');

--
-- Trigger `transaction_details`
--
DELIMITER $$
CREATE TRIGGER `updateTransQty` AFTER DELETE ON `transaction_details` FOR EACH ROW BEGIN
DECLARE total_qty_new INT;
SET total_qty_new = (SELECT total_qty FROM transactions WHERE id=OLD.transaction_id);
UPDATE transactions SET total_qty = total_qty_new - OLD.qty WHERE id= OLD.transaction_id;
END
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `status` smallint(11) NOT NULL,
  `dob` date NOT NULL,
  `gender` char(1) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `status`, `dob`, `gender`, `created_at`, `updated_at`) VALUES
(1, 'yanto', 1, '2001-03-04', 'L', '2023-03-17 04:50:18', '2023-03-17 04:50:18'),
(2, 'slamet', 1, '2001-03-12', 'P', '2023-03-17 04:50:27', '2023-03-17 04:50:27'),
(3, 'john doe', 1, '2001-07-22', 'P', '2023-03-17 04:50:34', '2023-03-17 04:50:34'),
(4, 'johan', 1, '2001-01-12', 'P', '2023-03-17 04:50:41', '2023-03-17 04:50:41'),
(5, 'putra', 1, '2001-06-27', 'L', '2023-03-17 04:50:49', '2023-03-17 04:50:49'),
(8, 'test aja', 1, '2023-03-15', 'P', '2023-03-01 22:06:53', '2023-03-17 16:06:53');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `operators`
--
ALTER TABLE `operators`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `payment_methods`
--
ALTER TABLE `payment_methods`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `operator_id` (`operator_id`),
  ADD KEY `product_type_id` (`product_type_id`);

--
-- Indeks untuk tabel `product_descriptions`
--
ALTER TABLE `product_descriptions`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `product_types`
--
ALTER TABLE `product_types`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `payment_method_id` (`payment_method_id`);

--
-- Indeks untuk tabel `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD PRIMARY KEY (`id`),
  ADD KEY `transaction_id` (`transaction_id`),
  ADD KEY `product_id` (`product_id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `operators`
--
ALTER TABLE `operators`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT untuk tabel `payment_methods`
--
ALTER TABLE `payment_methods`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT untuk tabel `product_descriptions`
--
ALTER TABLE `product_descriptions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT untuk tabel `product_types`
--
ALTER TABLE `product_types`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT untuk tabel `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- AUTO_INCREMENT untuk tabel `transaction_details`
--
ALTER TABLE `transaction_details`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=46;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_3` FOREIGN KEY (`operator_id`) REFERENCES `operators` (`id`),
  ADD CONSTRAINT `products_ibfk_4` FOREIGN KEY (`product_type_id`) REFERENCES `product_types` (`id`);

--
-- Ketidakleluasaan untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `transactions_ibfk_2` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_methods` (`id`);

--
-- Ketidakleluasaan untuk tabel `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD CONSTRAINT `transaction_details_ibfk_1` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`),
  ADD CONSTRAINT `transaction_details_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
