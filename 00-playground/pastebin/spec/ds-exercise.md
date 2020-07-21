# B"H



Intsights Exercise:

---



### Write a simple web crawler.

- Crawl the site: https://pastebin.com/ and should store the most recent "pastes" in a structured data model.

- A paste model must have the following parameters:
    - Author - String
    - Title - String
    - Content - String
    - Date - Date


The code must be self managed. It should crawl the site every 2 minutes and look for any new pastes to save.


--- --- --- --- --- --- --- --- --- --- ---

### Bonus #1

Each one of the paste model's parameters must be normalized.

For example:
- Author 
    - In cases it's Guest, Unknown, Anonymous, etc... the author name must be the same, for example: "" (empty string)
    - Title - Same as with Author.
    - Date - UTC Date
    - Content - Must be stripped of trailing spaces.


### Bonus #2
Store each one of the pastes in an organized database. 


### Bonus #3
Ship your crawler in a Docker image. This will allow it to run on any platform and any computer without any necessary installations. docker-compose and docker swarm solutions are also welcome.


### Bonus #4 - Data Analysis

First, please run your crawler for a several hours (up to 1-2 days) to generate some data. This is the data you'll be working on. 

The task here is to clusterize the data that you've found to at least 3 groups (can be more). The groups are yours to decide on, but you can use one or more of the groups below. 

1. Credentials Leakage
2. Source Code
3. Technical Data
4. Encoded binary data (base64 or others)
5. List of URLs



Please include all the relevant code you used for clusterization and classification. Feel free to include jupyter notebooks, and any other resources that you used for the analysis process.



General Notes (Apply with or without the bonus parts):
------------------------------------------------------
The code must be supplied with a README.md that explains how to create the environment for the code to run.


Emphasis should be placed on the following:
    - Write clean code. Do not document unnecessary lines, (you can consult the "Zen of Python" - https://gist.github.com/evandrix/2030615) 
    - Use Object Oriented Methodologies
    - The code must be readable, and the directory tree must be chosen wisely.
    - The code must not have any flake8 (linting) warnings at all.

We strongly recommend using the following libraries:
    - requests (http client)
    - lxml (html parsing)
    - tinydb (storage)
    - arrow (dates)

