## Pagination

### Front End

- Add `Paginator` Component
  - Keeps current page number as state
  - On Prev/Next button click -> trigger an API request
    - encode page number and page size as query parameters
      - ex: `'/api/tracks?page=2&pageSize=20'`
- disable Prev/Next buttons according to either pagination metadata or checking client-side values (`if data.length < PAGE_SIZE`, it's the last page)

### Back End

- In route handler, parse out query params, read values into the struct, have appropriate defaults
- Implement functions on the struct to calculate/return `limit` and `offset`
- Pass in limit and offset to model function
- Modify SQL to use limit and offset

  - ex:

  ```
   SELECT t.id, t.name, a.title, mt.name, g.name, composer, milliseconds, bytes, unit_price
  	FROM tracks t
  	INNER JOIN albums a ON a.id = t.album_id
  	INNER JOIN media_types mt ON mt.id = t.media_type_id
  	INNER JOIN genres g ON g.id = t.genre_id
  	LIMIT $1 OFFSET $2

  ```

## Search

### Front End

- Create Search component with a controlled text input, a submit button, and a reset button
- on submit, send the search input state to the api as a query parameter, clear the text input
- keep the current search term as state in the parent Table component so that when a term returns
  multiple pages, the "Next Page" data is still associated with the specific search term

### Back End

- parse out the 'query' query parameter into the same struct as the pagination inputs
- pass in the query to the model and use ILIKE or full text search
- ex:
  ```
   SELECT t.id, t.name, a.title, mt.name, g.name, composer, milliseconds, bytes, unit_price
  	FROM tracks t
  	INNER JOIN albums a ON a.id = t.album_id
  	INNER JOIN media_types mt ON mt.id = t.media_type_id
  	INNER JOIN genres g ON g.id = t.genre_id
    WHERE t.name ILIKE $1
  	LIMIT $2 OFFSET $3
  ```

## Serve Static React App Build from Go Router

- Create an http.FileServer and pass in the path to the `ui/build`
- Create a home route `/` that uses the fileServer as a Handler and let React Router handle the client side routing accordingly
