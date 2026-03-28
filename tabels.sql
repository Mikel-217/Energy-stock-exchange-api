
CREATE TABLE EnergyPrice(
    EnergyPriceId INT UNSIGNED NOT NULL AUTO_INCREMENT,
    CurrentDate DATETIME,
    BestTimeToBuy INT UNSIGNED, -- foreign key
    BestTimeToSell INT UNSIGNED, -- foreign key

    PRIMARY KEY(EnergyPriceId)
);

CREATE TABLE DateAndPrice(
    DatePriceId INT UNSIGNED NOT NULL AUTO_INCREMENT,
    Date DATETIME,
    Price FLOAT,
    EnergyPriceId INT UNSIGNED, -- foreign key

    FOREIGN KEY(EnergyPriceId) REFERENCES EnergyPrice(EnergyPriceId),
    PRIMARY KEY(DatePriceId)
);

ALTER TABLE EnergyPrice
ADD CONSTRAINT fk_best_buy
FOREIGN KEY (BestTimeToBuy) REFERENCES DateAndPrice(DatePriceId);

ALTER TABLE EnergyPrice
ADD CONSTRAINT fk_best_sell
FOREIGN KEY (BestTimeToSell) REFERENCES DateAndPrice(DatePriceId);
