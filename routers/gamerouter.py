from core.templates import templates
from fastapi import APIRouter, HTTPException, Query, Request
from fastapi.responses import HTMLResponse, JSONResponse
from steamapi import SteamAPI
from utils import BannerHelper
from pydantic import BaseModel

client = SteamAPI()

game_router = APIRouter()

banner_helper = BannerHelper()

class Body(BaseModel):
    game_ids: list[int]

@game_router.get("/", response_class=HTMLResponse)
async def get_header_by_game_id(
        request: Request,
        body: Body, 
        bannertype: str = Query(..., regex="^(boxart|header)$"),
    ):

    try:
        if bannertype == "header":
            banners = banner_helper.get_header_urls(body.game_ids)
        else:
            banners = banner_helper.get_box_art_urls(body.game_ids)
        
        if request.headers.get("hx-request") == "true":
            return templates.TemplateResponse("partials/_game_banners.html", { "request": request, "games": body.game_ids })
    
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error fetching banners: {str(e)}")
    
    return JSONResponse(banners)