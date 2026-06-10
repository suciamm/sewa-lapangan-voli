CREATE TABLE owner_verifications (
    id          BIGINT       NOT NULL AUTO_INCREMENT,
    user_id     BIGINT       NOT NULL,
    ktp_url     VARCHAR(500),
    doc_url     VARCHAR(500),
    status      ENUM('pending','approved','rejected') NOT NULL DEFAULT 'pending',
    reviewed_by BIGINT,
    reviewed_at DATETIME,
    notes       TEXT,
    created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY idx_ov_user_id     (user_id),
    KEY idx_ov_reviewed_by (reviewed_by),
    CONSTRAINT fk_ov_user     FOREIGN KEY (user_id)     REFERENCES users (id),
    CONSTRAINT fk_ov_reviewer FOREIGN KEY (reviewed_by) REFERENCES users (id)
);
