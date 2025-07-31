from fastapi import APIRouter
from steamapi import SteamAPI
from utils import BannerHelper

client = SteamAPI()

game_router = APIRouter()