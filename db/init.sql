CREATE TABLE Usuario (
    ID INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    Nome VARCHAR(50),
    Tipo VARCHAR(50),
    Cpf VARCHAR(50),
    Email VARCHAR(50),
    Saldo FLOAT(10.2)
);

CREATE TABLE Transacao (
    ID INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    IDPayer INT NOT NULL,
    IDPayee INT NOT NULL,
    Valor FLOAT(10.2)
);

"INSERT INTO tbnecessidade "+"(desc_necessidade, progresso_necessidade, fk_id_material, fk_id_tipo)"
        + " VALUES (?, ?, (select max(pk_id_Material) from tbmaterial),(select max(pk_id_Tipo) from tbTipoOperacao))"

INSERT INTO usuario (Nome, Saldo)VALUES ("Izaque", 100);