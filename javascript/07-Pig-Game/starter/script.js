'use strict';

// DOM Elements
const dice = document.querySelector('.dice');
const player1 = document.querySelector('.player--0');
const player2 = document.querySelector('.player--1');

// Clickable Elements
const newGameBtn = document.querySelector('.btn--new');
const btnRoll = document.querySelector('.btn--roll');
const btnhold = document.querySelector('.btn--hold');
// Score Display Elements
const player1ScoreEl = document.querySelector('#score--0');
const player2ScoreEl = document.querySelector('#score--1');
const player1CurrentEl = document.querySelector('#current--0');
const player2CurrentEl = document.querySelector('#current--1');
// Game Tracking Variables
let player1Score = 0;
let player2Score = 0;
let player1Current = 0;
let player2Current = 0;
let activePlayer = 1;
let playing = true;
// Refesh UI
updateUI();

function rollDice() {
  if (!playing) return;

  const randomNumber = Math.floor(Math.random() * 6) + 1;
  dice.src = `dice-${randomNumber}.png`;

  if (randomNumber == 1) {
    switchPlayer();
  } else {
    activePlayer == 1
      ? (player1Current += randomNumber)
      : (player2Current += randomNumber);
  }
  updateUI();
}

function switchPlayer() {
  if (activePlayer == 1) {
    player1Current = 0;
    activePlayer = 2;
    player1.classList.remove('player--active');
    player2.classList.add('player--active');
  } else {
    player2Current = 0;
    activePlayer = 1;
    player2.classList.remove('player--active');
    player1.classList.add('player--active');
  }
  updateUI();
}

function hold() {
  if (!playing) return;

  if (activePlayer == 1) {
    player1Score += player1Current;
  } else {
    player2Score += player2Current;
  }

  if (player1Score >= 100) {
    player1.classList.add('player--winner');
    alert('Player 1 Wins!');
    playing = false;
  }

  if (player2Score >= 100) {
    player2.classList.add('player--winner');
    alert('Player 2 Wins!');
    playing = false;
  }

  switchPlayer();
  updateUI();
}

function newGame() {
  player1Score = 0;
  player2Score = 0;
  player1Current = 0;
  player2Current = 0;
  activePlayer = 1;

  player2.classList.remove('player--active');
  player1.classList.add('player--active');

  player2.classList.remove('player--winner');
  player1.classList.remove('player--winner');
  updateUI();
}

function updateUI() {
  player1ScoreEl.textContent = player1Score;
  player2ScoreEl.textContent = player2Score;
  player1CurrentEl.textContent = player1Current;
  player2CurrentEl.textContent = player2Current;
}

// Add Event Listeners
newGameBtn.addEventListener('click', newGame);
btnRoll.addEventListener('click', rollDice);
btnhold.addEventListener('click', hold);
