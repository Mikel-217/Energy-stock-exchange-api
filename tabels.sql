
CREATE TABLE EnergyPrice(
    EnergyPriceId INT UNSIGNED NOT NULL AUTO_INCREMENT,
    CurrentDate DATETIME,
    BestTimeToBuy INT UNSIGNED NOT NULL, -- foreign key
    BestTimeToSell INT UNSIGNED NOT NULL, -- foreign key

    FOREIGN KEY(BestTimeToBuy) REFERENCES DateAndPrice(DatePriceId),
    FOREIGN KEY(BestTimeToSell) REFERENCES DateAndPrice(DatePriceId),
    PRIMARY KEY(EnergyPriceId)
);

CREATE TABLE DateAndPrice(
    DatePriceId INT UNSIGNED NOT NULL AUTO_INCREMENT,
    Date DATETIME,
    Price FLOAT,
    EnergyPriceId INT UNSIGNED NOT NULL, -- foreign key

    FOREIGN KEY(EnergyPriceId) REFERENCES EnergyPrice(EnergyPriceId),
    PRIMARY KEY(DatePriceId)
);
