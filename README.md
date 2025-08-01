# Steam Banner API

- [Description](#description)
- [Usage](#usage)
    - [Query Parameters](#query-parameters)
    - [Search by Game IDs](#search-by-game-ids)
    - [Search by User ID](#search-by-user-id)

## Description
An API for obtaining Steam game banners with custom search filters. This API exists to work around the limitations of the Steam Web API allowing more advanced filtering for obtaining game image banners. 

## Usage

### Search by Game IDs
- `/api/game?...`

### Search by User ID
- `/api/profile?...`

### Query Parameters
- `bannertype` (required)
    - `box-art`
        - Payload returns game box art.
    - `header`
        - Payload returns game header.
- `completed` (required for profile filtering)
    - `true`
        - Payload returns games completed by user.
    - `false`
        - Payload returns all games regardless of completion status.
