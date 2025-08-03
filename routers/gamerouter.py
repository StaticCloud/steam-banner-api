from fastapi import APIRouter, HTTPException, Query
from steamapi import SteamAPI
from utils import BannerHelper
from pydantic import BaseModel

client = SteamAPI()

game_router = APIRouter()

banner_helper = BannerHelper()

class Body(BaseModel):
    game_ids: list[int]

@game_router.get("/")
def get_header_by_game_id(
        body: Body, 
        bannertype: str = Query(..., regex="^(boxart|header)$")
    ):

    try:
        if bannertype == "header":
            banners = banner_helper.get_header_urls(body.game_ids)
        else:
            banners = banner_helper.get_box_art_urls(body.game_ids)
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error fetching banners: {str(e)}")
    
    return banners