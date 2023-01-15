from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from os import path



def create_app():
    print("creating app")
    app = Flask(__name__)
    app.config['SECRET_KEY'] = 'hjshjhdkjshkjdhjss'
    # app.config['SQLALCHEMY_DATABASE_URI'] = r"sqlite:///./db.sqlite"
create_app()