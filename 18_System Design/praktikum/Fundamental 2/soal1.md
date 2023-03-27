Terdapat sebuah query pada SQL yaitu `SELECT * FROM users;`
Dengan tujuan yang sama, tuliskan dalam bentuk perintah:

<ol>
<li>Redis</li>
	
```KEYS *Users*```
	
<li>Neo4j</li>
	
```MATCH (u:Users) RETURN u```
	
<li>Cassandra</li>
	
```SELECT * FROM keySpace.Users;```
	
</ol>