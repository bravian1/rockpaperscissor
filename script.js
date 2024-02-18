document.addEventListener('DOMContentLoaded', function() {
    const rockBtn = document.getElementById('rock');
    const paperBtn = document.getElementById('paper');
    const scissorsBtn = document.getElementById('scissors');
    const playerScoreSpan = document.getElementById('player-score');
    const computerScoreSpan = document.getElementById('computer-score');
    const resultMessage = document.getElementById('result-message');
    const computerChoiceSpan = document.getElementById('computer-choice'); // Reference to the new element

    // Event Listeners for Button Clicks
    rockBtn.addEventListener('click', () => sendMoveToServer('rock'));
    paperBtn.addEventListener('click', () => sendMoveToServer('paper'));
    scissorsBtn.addEventListener('click', () => sendMoveToServer('scissors'));

    // Helper function to send player's move to the Go server
    function sendMoveToServer(move) {
        fetch('http://localhost:8080/rps', {  // Updated to include full URL with port
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded' // Specify data type 
            },
            body: 'move=' + move  // Send the player's move as form data
        })
        .then(response => response.json())  // Expect JSON response from server
        .then(gameResult => {
            updateUI(gameResult); 
        })
        .catch(error => console.error('Error:', error));  // Error handling
    }

    // Function to update the UI based on the game result received from the server
    function updateUI(gameResult) {
        playerScoreSpan.textContent = gameResult.playerScore;
        computerScoreSpan.textContent = gameResult.computerScore;
        resultMessage.textContent = gameResult.message; // Display the result message
        computerChoiceSpan.textContent = gameResult.computerChoice; // Display the computer's choice
    }
});