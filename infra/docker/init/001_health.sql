-- DB 接続の確認ポイントとして、「繋がって、１件だけ取得できる」単純なテーブルを用意

CREATE TABLE health_check (
  id INT PRIMARY KEY AUTO_INCREMENT,
  message VARCHAR(255) NOT NULL
);

INSERT INTO health_check (message) VALUES ('ok');
