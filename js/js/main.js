var board = null
var game = new Chess()
var $status = $('#status')
var $fen = $('#fen')
var $pgn = $('#pgn')
var ws = null;

function onDragStart (source, piece, position, orientation) {
  // do not pick up pieces if the game is over
  if (game.game_over()) return false

  // only pick up pieces for the side to move
  if ((game.turn() === 'w' && piece.search(/^b/) !== -1) ||
      (game.turn() === 'b' && piece.search(/^w/) !== -1)) {
    return false
  }
}

function onDrop (source, target) {
  // see if the move is legal
  var move = game.move({
    from: source,
    to: target,
    promotion: 'q' // TODO: always promote to a queen for example simplicity
  })

  // illegal move
  if (move === null) return 'snapback'

  updateStatus()
}

// update the board position after the piece snap
// for castling, en passant, pawn promotion
function onSnapEnd () {
  let fen = game.fen();
  board.position(fen)
  msg = "position " + fen;
  console.log(msg);
  ws.send(msg);

  msg = "go infinite";
  console.log(msg);
  ws.send(msg);
}

function updateStatus () {
  var status = ''

  var moveColor = 'White'
  if (game.turn() === 'b') {
    moveColor = 'Black'
  }

  // checkmate?
  if (game.in_checkmate()) {
    status = 'Game over, ' + moveColor + ' is in checkmate.'
  }

  // draw?
  else if (game.in_draw()) {
    status = 'Game over, drawn position'
  }

  // game still on
  else {
    status = moveColor + ' to move'

    // check?
    if (game.in_check()) {
      status += ', ' + moveColor + ' is in check'
    }
  }

  $status.html(status)
  $fen.html(game.fen())
  $pgn.html(game.pgn())
}

function init() {
    var config = {
    draggable: true,
    position: 'start',
    onDragStart: onDragStart,
    onDrop: onDrop,
    onSnapEnd: onSnapEnd
    }

    ws = new WebSocket("ws://" + document.location.host + "/ws");
    ws.onmessage = function (evt) {
        console.log(evt.data);
        let split = evt.data.split(" ");
        if (split[0] === "bestmove") {
          let move = split[1];
          let from = move.substr(0,2);
          let to = move.substr(2,2);

          // see if the move is legal
          if (game.move({
            from: from,
            to: to,
            promotion: 'q' // TODO: always promote to a queen for example simplicity
          }) === null) {
            return 'snapback';
          }

          board.position(game.fen());

          updateStatus();
        }
    };

    board = Chessboard('myBoard', config)

    updateStatus()

    document.getElementById("newGame").onclick = function () {
        board = ChessBoard('myBoard', config);
        msg = "ucinewgame";
        console.log(msg);
        ws.send(msg);
        return false;
    };
}

$(document).ready(init);