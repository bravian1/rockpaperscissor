document.addEventListener('DOMContentLoaded', function() {
    const rockBtn = document.getElementById('rock');
    const paperBtn = document.getElementById('paper');
    const scissorsBtn = document.getElementById('scissors');
    const playerScoreSpan = document.getElementById('player-score');
    const computerScoreSpan = document.getElementById('computer-score');
    const resultMessage = document.getElementById('result-message');
    const computerChoiceSpan = document.getElementById('computer-choice');
    const historyTable = document.getElementById('history-table');
    const newSetBtn = document.getElementById('new-set');
    
    let targetWins; // Number of wins needed for a set
    let gameHistory = []; // Array to keep track of game history

    // Function to reset scores on the server
    function resetServerScores() {
        fetch('http://localhost:8080/reset', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
            }
        })
        .then(response => response.json())
        .then(data => console.log(data.message))
        .catch(error => console.error('Error:', error));
    }

    // Function to start a new game set
    function startNewSet() {
        targetWins = parseInt(prompt("How many wins are needed to win a set?"));
        if (isNaN(targetWins) || targetWins <= 0) {
            alert("Please enter a valid number of wins.");
            return;
        }
        resetServerScores(); // Reset scores on the server
        playerScoreSpan.textContent = '0';
        computerScoreSpan.textContent = '0';
        resultMessage.textContent = '';
        computerChoiceSpan.textContent = '';
        gameHistory = []; // Reset game history
        updateHistoryTable(); // Update the history table
    }

    // Function to update the game history table
    function updateHistoryTable() {
        historyTable.innerHTML = ''; // Clear the table
        gameHistory.forEach((result, index) => {
            const row = historyTable.insertRow(index);
            const cell1 = row.insertCell(0);
            const cell2 = row.insertCell(1);
            const cell3 = row.insertCell(2);
            cell1.textContent = `Game ${index + 1}`;
            cell2.textContent = result.playerChoice;
            cell3.textContent = result.computerChoice;
        });
    }

    // Function to check for set winner and prompt for new game
    function checkForSetWinner(playerScore, computerScore) {
        if (playerScore >= targetWins || computerScore >= targetWins) {
            const winner = playerScore > computerScore ? "Player" : "Computer";
            alert(`${winner} wins the set!`);
            if (confirm("Would you like to play again?")) {
                startNewSet();
            } else {
                rockBtn.disabled = true;
                paperBtn.disabled = true;
                scissorsBtn.disabled = true;
            }
        }
    }

    // Event Listeners for Button Clicks
    rockBtn.addEventListener('click', () => sendMoveToServer('rock'));
    paperBtn.addEventListener('click', () => sendMoveToServer('paper'));
    scissorsBtn.addEventListener('click', () => sendMoveToServer('scissors'));
    // Event listener for the new set button
    newSetBtn.addEventListener('click', startNewSet);

    // Helper function to send player's move to the Go server
    function sendMoveToServer(move) {
        fetch('http://localhost:8080/rps', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
            },
            body: 'move=' + move
        })
        .then(response => response.json())
        .then(gameResult => {
            updateUI(gameResult);
            gameHistory.push({ playerChoice: move, computerChoice: gameResult.computerChoice });
            updateHistoryTable();
            checkForSetWinner(gameResult.playerScore, gameResult.computerScore);
        })
        .catch(error => console.error('Error:', error));
    }

    // Function to update the UI based on the game result received from the server
    function updateUI(gameResult) {
        playerScoreSpan.textContent = gameResult.playerScore;
        computerScoreSpan.textContent = gameResult.computerScore;
        resultMessage.textContent = gameResult.message;
        computerChoiceSpan.textContent = gameResult.computerChoice;
    }

    // Start the first set
    startNewSet();
});