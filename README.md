Should I move the folders into adapters folder to the root folder?

Como lidar com WebSocket:

1 alternativa: Ao solicitar o endpoint de websocket enviando uma pollID, retornar
os resultados dessa pollID e depois retornar os novos resultados dessa poll via WS sempre que
recer uma nova mensagem de novo voto no canal dessa poll

2 alternativa: Ao solicitar o endpoint de websocket enviando uma pollID, retornar
os resultados dessa pollID e depois enviar uma nova mensagem via WS notificando que um novo voto foi computado
sempre que receber uma nova mensagem de novo voto no canal dessa poll, assim o cliente ficaria responsavel por fazer
uma nova request a uma rota http normal sempre que quiser (a cada 1 voto, a cada 5 votos, etc...)


em ambos os casos eu precisaria usar um channel de votação para cada pollID (exemplo: new_vote_in_7s5da21wasd23555a)
e precisaria talvez tambem guardar em memória cada uma das conexões