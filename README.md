Listar usuarios:

		# curl -i http://localhost:3113/api/usuarios

Mostrar un solo usuario (reemplazar {id} por bson.ObjectIdHex):

		# curl -i http://localhost:3113/api/usuarios/{id}

Registrar usuarios:

		# curl -i http://localhost:3113/api/usuarios -X POST -d @add.json

Modificar un usuario (reemplazar {id} por bson.ObjectIdHex):

		# curl -i http://localhost:3113/api/usuarios/{id} -X PUT -d @update.json

Borrar un usuario:

		# curl -i http://localhost:3113/api/usuarios/{id} -X DELETE

Buscar un usuario por el campo usuario (reemplazar {usuario} por usuario a buscar):

		# curl -i http://localhost:3113/api/usuarios/buscar/porusuario/{usuario}
