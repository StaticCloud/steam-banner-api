from concurrent.futures import ThreadPoolExecutor, as_completed
from dotenv import load_dotenv
import requests
import os

load_dotenv()

class SteamAPI():
    def __init__(self):
        self.api_key = os.getenv("API_KEY")

    def get_owned_game_ids(self, steam_id: str):
        try:
            url = f'http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key={self.api_key}&steamid={steam_id}&format=json'
            res = requests.get(url, 5).json()

            print(res)

            game_ids = [game['appid'] for game in res['response']['games']]

            return game_ids
        except Exception as e:
            return []
    
    def get_completed_game_ids(self, steam_id: str):
        game_ids = self.get_owned_game_ids(steam_id)

        completed_game_ids = []

        def check_completion(game_id):
            try:
                url = f'http://api.steampowered.com/ISteamUserStats/GetPlayerAchievements/v0001/?appid={game_id}&key={self.api_key}&steamid={steam_id}'
                res = requests.get(url, 5).json()

                achievements = res['playerstats']['achievements']

                if achievements and self.validate_completion_status(achievements):
                    return game_id
            except Exception as e:
                pass
            return None
        
        with ThreadPoolExecutor(10) as executor:
            futures = [executor.submit(check_completion, game_id) for game_id in game_ids]

            for future in as_completed(futures):
                result = future.result()

                if result:
                    completed_game_ids.append(result)
            
        return completed_game_ids 

    def validate_completion_status(self, achievements: list):
            return all(achievement['achieved'] == 1 for achievement in achievements)
    