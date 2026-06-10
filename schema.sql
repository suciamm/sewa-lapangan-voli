-- ─────────────────────────────────────────
-- USERS & AUTH
-- ─────────────────────────────────────────

CREATE TABLE users (
    id         BIGINT       NOT NULL AUTO_INCREMENT,
    name       VARCHAR(100) NOT NULL,
    email      VARCHAR(150) NOT NULL,
    password   VARCHAR(255) NOT NULL,
    phone      VARCHAR(20),
    role       ENUM('superuser','owner','penyewa') NOT NULL,
    status     ENUM('active','inactive','pending') NOT NULL DEFAULT 'pending',
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY uq_users_email (email)
);

CREATE TABLE owner_verifications (
    id          BIGINT    NOT NULL AUTO_INCREMENT,
    user_id     BIGINT    NOT NULL,
    ktp_url     VARCHAR(500),
    doc_url     VARCHAR(500),
    status      ENUM('pending','approved','rejected') NOT NULL DEFAULT 'pending',
    reviewed_by BIGINT,
    reviewed_at DATETIME,
    notes       TEXT,
    created_at  DATETIME  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY idx_owner_verifications_user_id (user_id),
    KEY idx_owner_verifications_reviewed_by (reviewed_by),
    CONSTRAINT fk_ov_user       FOREIGN KEY (user_id)     REFERENCES users (id),
    CONSTRAINT fk_ov_reviewer   FOREIGN KEY (reviewed_by) REFERENCES users (id)
);

CREATE TABLE refresh_tokens (
    id         BIGINT       NOT NULL AUTO_INCREMENT,
    user_id    BIGINT       NOT NULL,
    token      VARCHAR(500) NOT NULL,
    expires_at DATETIME     NOT NULL,
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY uq_refresh_tokens_token (token),
    KEY idx_refresh_tokens_user_id (user_id),
    CONSTRAINT fk_rt_user FOREIGN KEY (user_id) REFERENCES users (id)
);

-- ─────────────────────────────────────────
-- LAPANGAN (COURTS)
-- ─────────────────────────────────────────

CREATE TABLE courts (
    id            BIGINT          NOT NULL AUTO_INCREMENT,
    owner_id      BIGINT          NOT NULL,
    name          VARCHAR(150)    NOT NULL,
    description   TEXT,
    address       TEXT            NOT NULL,
    province      VARCHAR(100)    NOT NULL,
    city          VARCHAR(100)    NOT NULL,
    status        ENUM('active','inactive','maintenance') NOT NULL DEFAULT 'active',
    avg_rating    DECIMAL(3,2)    NOT NULL DEFAULT 0.00,
    total_reviews INT             NOT NULL DEFAULT 0,
    created_at    DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY idx_courts_owner_id (owner_id),
    KEY idx_courts_province (province),
    CONSTRAINT fk_courts_owner FOREIGN KEY (owner_id) REFERENCES users (id)
);

CREATE TABLE court_prices (
    id             BIGINT         NOT NULL AUTO_INCREMENT,
    court_id       BIGINT         NOT NULL,
    day_type       ENUM('weekday','weekend') NOT NULL,
    session        ENUM('morning','afternoon','night') NOT NULL,
    price_per_hour DECIMAL(12,2)  NOT NULL,
    PRIMARY KEY (id),
    KEY idx_court_prices_court_id (court_id),
    CONSTRAINT fk_cp_court FOREIGN KEY (court_id) REFERENCES courts (id)
);

CREATE TABLE court_images (
    id         BIGINT       NOT NULL AUTO_INCREMENT,
    court_id   BIGINT       NOT NULL,
    image_url  VARCHAR(500) NOT NULL,
    is_primary TINYINT(1)   NOT NULL DEFAULT 0,
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY idx_court_images_court_id (court_id),
    CONSTRAINT fk_ci_court FOREIGN KEY (court_id) REFERENCES courts (id)
);

-- ─────────────────────────────────────────
-- BOOKING & RESCHEDULE
-- ─────────────────────────────────────────

CREATE TABLE bookings (
    id             BIGINT        NOT NULL AUTO_INCREMENT,
    court_id       BIGINT        NOT NULL,
    penyewa_id     BIGINT        NOT NULL,
    booking_date   DATE          NOT NULL,
    start_time     TIME          NOT NULL,
    end_time       TIME          NOT NULL,
    duration_hours INT           NOT NULL,
    total_price    DECIMAL(12,2) NOT NULL,
    status         ENUM('pending_payment','paid','active','completed','reschedule_pending') NOT NULL DEFAULT 'pending_payment',
    hold_until     DATETIME,
    created_at     DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY idx_bookings_court_id (court_id),
    KEY idx_bookings_penyewa_id (penyewa_id),
    KEY idx_bookings_booking_date (booking_date),
    CONSTRAINT fk_bookings_court   FOREIGN KEY (court_id)   REFERENCES courts (id),
    CONSTRAINT fk_bookings_penyewa FOREIGN KEY (penyewa_id) REFERENCES users  (id)
);

CREATE TABLE reschedule_requests (
    id             BIGINT    NOT NULL AUTO_INCREMENT,
    booking_id     BIGINT    NOT NULL,
    requested_by   BIGINT    NOT NULL,
    new_date       DATE      NOT NULL,
    new_start_time TIME      NOT NULL,
    new_end_time   TIME      NOT NULL,
    status         ENUM('pending','approved','rejected','expired') NOT NULL DEFAULT 'pending',
    reason         TEXT,
    reviewed_by    BIGINT,
    reviewed_at    DATETIME,
    hold_until     DATETIME,
    created_at     DATETIME  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     DATETIME  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY idx_rr_booking_id   (booking_id),
    KEY idx_rr_requested_by (requested_by),
    KEY idx_rr_reviewed_by  (reviewed_by),
    CONSTRAINT fk_rr_booking      FOREIGN KEY (booking_id)   REFERENCES bookings (id),
    CONSTRAINT fk_rr_requester    FOREIGN KEY (requested_by) REFERENCES users    (id),
    CONSTRAINT fk_rr_reviewer     FOREIGN KEY (reviewed_by)  REFERENCES users    (id)
);

CREATE TABLE court_reviews (
    id           BIGINT    NOT NULL AUTO_INCREMENT,
    booking_id   BIGINT    NOT NULL,
    court_id     BIGINT    NOT NULL,
    penyewa_id   BIGINT    NOT NULL,
    rating       TINYINT   NOT NULL,  -- 1–5
    comment      TEXT,
    owner_reply  TEXT,
    replied_at   DATETIME,
    created_at   DATETIME  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY uq_court_reviews_booking_id (booking_id),
    KEY idx_court_reviews_court_id   (court_id),
    KEY idx_court_reviews_penyewa_id (penyewa_id),
    CONSTRAINT fk_cr_booking FOREIGN KEY (booking_id) REFERENCES bookings (id),
    CONSTRAINT fk_cr_court   FOREIGN KEY (court_id)   REFERENCES courts   (id),
    CONSTRAINT fk_cr_penyewa FOREIGN KEY (penyewa_id) REFERENCES users    (id)
);

-- ─────────────────────────────────────────
-- PEMBAYARAN & KEUANGAN
-- ─────────────────────────────────────────

CREATE TABLE payments (
    id                BIGINT        NOT NULL AUTO_INCREMENT,
    booking_id        BIGINT        NOT NULL,
    midtrans_order_id VARCHAR(100)  NOT NULL,
    midtrans_tx_id    VARCHAR(100),
    amount            DECIMAL(12,2) NOT NULL,
    payment_method    VARCHAR(50),
    status            ENUM('pending','settlement','expire','failure') NOT NULL DEFAULT 'pending',
    snap_token        VARCHAR(500),
    paid_at           DATETIME,
    created_at        DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY uq_payments_booking_id        (booking_id),
    UNIQUE KEY uq_payments_midtrans_order_id (midtrans_order_id),
    CONSTRAINT fk_payments_booking FOREIGN KEY (booking_id) REFERENCES bookings (id)
);

CREATE TABLE platform_fees (
    id           BIGINT        NOT NULL AUTO_INCREMENT,
    payment_id   BIGINT        NOT NULL,
    court_id     BIGINT        NOT NULL,
    owner_id     BIGINT        NOT NULL,
    gross_amount DECIMAL(12,2) NOT NULL,
    fee_percent  DECIMAL(5,2)  NOT NULL,
    fee_amount   DECIMAL(12,2) NOT NULL,
    net_to_owner DECIMAL(12,2) NOT NULL,
    created_at   DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY uq_platform_fees_payment_id (payment_id),
    KEY idx_platform_fees_court_id (court_id),
    KEY idx_platform_fees_owner_id (owner_id),
    CONSTRAINT fk_pf_payment FOREIGN KEY (payment_id) REFERENCES payments (id),
    CONSTRAINT fk_pf_court   FOREIGN KEY (court_id)   REFERENCES courts   (id),
    CONSTRAINT fk_pf_owner   FOREIGN KEY (owner_id)   REFERENCES users    (id)
);

CREATE TABLE platform_fee_config (
    id             BIGINT       NOT NULL AUTO_INCREMENT,
    fee_percent    DECIMAL(5,2) NOT NULL,
    effective_from DATE         NOT NULL,
    created_by     BIGINT       NOT NULL,
    created_at     DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY idx_pfc_effective_from (effective_from),
    KEY idx_pfc_created_by     (created_by),
    CONSTRAINT fk_pfc_creator FOREIGN KEY (created_by) REFERENCES users (id)
);

-- ─────────────────────────────────────────
-- NOTIFIKASI
-- ─────────────────────────────────────────

CREATE TABLE notifications (
    id         BIGINT       NOT NULL AUTO_INCREMENT,
    user_id    BIGINT       NOT NULL,
    type       ENUM('booking_confirmed','reschedule_request','reschedule_approved','reschedule_rejected','payment_success','payment_failed') NOT NULL,
    title      VARCHAR(200) NOT NULL,
    body       TEXT         NOT NULL,
    ref_id     BIGINT,
    is_read    TINYINT(1)   NOT NULL DEFAULT 0,
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY idx_notifications_user_id (user_id),
    CONSTRAINT fk_notif_user FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE email_verifications (
    id         BIGINT       NOT NULL AUTO_INCREMENT,
    user_id    BIGINT       NOT NULL,
    token      VARCHAR(100) NOT NULL,
    expires_at DATETIME     NOT NULL,
    used_at    DATETIME,
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY uq_ev_token (token),
    KEY idx_ev_user_id (user_id),
    CONSTRAINT fk_ev_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);