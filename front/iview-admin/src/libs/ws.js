import neffos from 'neffos.js'

// var scheme = document.location.protocol === 'https:' ? 'wss' : 'ws'
// var port = document.location.port ? ':' + document.location.port : ''

var wsURL = 'ws://127.0.0.1:8888/echo'

export const InitGlobalWebsocket = async () => {
  try {
    const conn = await neffos.dial(wsURL, {
      default: { // "default" namespace.
        _OnNamespaceConnected: function (nsConn, msg) {
          console.log('connected to namespace: ' + msg.Namespace)
          nsConn.emit('chat', 'Hello from browser(ify) client-side!')
        },
        _OnNamespaceDisconnect: function (nsConn, msg) {
          console.log('disconnected from namespace: ' + msg.Namespace)
        },
        chat: function (nsConn, msg) { // "chat" event.

        }
      }
    })
    // You can either wait to conenct or just conn.connect("connect")
    // and put the `handleNamespaceConnectedConn` inside `_OnNamespaceConnected` callback instead.
    // const nsConn = await conn.connect("default");
    // handleNamespaceConnectedConn(nsConn);
    // nsConn.emit(...); handleNamespaceConnectedConn(nsConn);
    console.log(neffos)
    conn.connect('default')
  } catch (err) {
    console.log(err)
  }
}


var outputTxt = document.getElementById("output")
function addMessage(msg) {
  outputTxt.innerHTML += msg + "\n"
}

function handleError(reason) {
  console.log(reason);
  window.alert(reason);
}

function handleNamespaceConnectedConn(nsConn) {
  nsConn.emit("chat", "Hello from browser(ify) client-side!");

  const inputTxt = document.getElementById("input");
  const sendBtn = document.getElementById("sendBtn");

  sendBtn.disabled = false;
  sendBtn.onclick = function () {
    const input = inputTxt.value;
    inputTxt.value = "";

    nsConn.emit("chat", input);
    addMessage("Me: " + input);
  };
}

export const runExample = async () => {
  try {
    const conn = await neffos.dial(wsURL, {
      default: { // "default" namespace.
        _OnNamespaceConnected: function (nsConn, msg) {
          addMessage("connected to namespace: " + msg.Namespace)
          handleNamespaceConnectedConn(nsConn);
        },
        _OnNamespaceDisconnect: function (nsConn, msg) {
          addMessage("disconnected from namespace: " + msg.Namespace)
        },
        chat: function (nsConn, msg) { // "chat" event.
          addMessage(msg.Body);
        }
      }
    })

    // You can either wait to conenct or just conn.connect("connect")
    // and put the `handleNamespaceConnectedConn` inside `_OnNamespaceConnected` callback instead.
    // const nsConn = await conn.connect("default");
    // handleNamespaceConnectedConn(nsConn);
    // nsConn.emit(...); handleNamespaceConnectedConn(nsConn);
    conn.connect("default")

  } catch (err) {
    handleError(err)
  }
}
