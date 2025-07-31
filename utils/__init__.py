class BannerHelper():
    def get_header_urls(self, game_ids):
        return [f"https://cdn.cloudflare.steamstatic.com/steam/apps/{game_id}/header.jpg" for game_id in game_ids]
    
    def get_box_art_urls(self, game_ids):
        return [f"https://cdn.steamstatic.com/steam/apps/{game_id}/library_600x900.jpg" for game_id in game_ids]