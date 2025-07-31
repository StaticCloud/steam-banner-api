# Steam Banner API

## Description
An API for obtaining Steam game banners with custom search filters. This API exists to work around the limitations of the Steam Web API allowing more advanced filtering for obtaining game image banners. 

## Usage

### Game IDs as Argument
- `/api/game/{banner_type}/{}`
    - `banner_type` Accepted Arguments:
        - `box-art`
        - `header`

### Profile IDs as Argument
- `/api/profile/{banner_type}/{steam_id}?completed=bool`
    - `banner_type` Accepted Arguments:
        - `box-art`
        - `header`
