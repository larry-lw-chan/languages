'use strict';
/* Use Prototype to create a object blueprint */
let Game = function (message, number, score, guess, again, check, highscore) {
  // Dom elements
  this.message = message;
  this.number = number;
  this.score = score;
  this.guess = guess;
  this.again = again;
  this.check = check;
  this.highScore = highscore;

  // Non-Dom attributes
  this.hiddenNumber = this.randomNumber();

  // Add Event Listener
  this.check.addEventListener('click', () => {
    this.checkGuess(this.guess.value);
  });

  this.again.addEventListener('click', () => {
    this.playAgain();
  });
};

Game.prototype.randomNumber = () => Math.trunc(Math.random() * 20) + 1;

Game.prototype.displayMessage = function (message) {
  this.message.textContent = message;
};

Game.prototype.updateScore = function (score) {
  this.score.textContent = score;
};

Game.prototype.playAgain = function (again) {
  this.hiddenNumber = this.randomNumber();
  this.score.textContent = '20';
  this.guess.value = '';
  this.number.textContent = '?';
  document.body.style.backgroundColor = '#222';
};

Game.prototype.checkGuess = function () {
  const guess = Number(this.guess.value);

  // Correct Guess Logic
  if (guess === this.hiddenNumber) {
    this.displayMessage('Correct Number!');
    this.number.textContent = this.hiddenNumber;
    document.body.style.backgroundColor = '#60b347';

    const score = Number(this.score.textContent);
    const highScore = Number(this.highScore.textContent);

    if (score > highScore) {
      this.highScore.textContent = score;
    }
    return;
  }

  // User entered no number
  if (!guess) {
    this.displayMessage('No Number!');
    return;
  }

  // Guess is too high or too low
  if (guess > this.hiddenNumber) {
    this.displayMessage('Too High!');
  } else if (guess < this.hiddenNumber) {
    this.displayMessage('Too Low!');
  }
  this.updateScore(this.score.textContent - 1);
};

/* Create Object Instance */
const game = new Game(
  document.querySelector('.message'),
  document.querySelector('.number'),
  document.querySelector('.score'),
  document.querySelector('.guess'),
  document.querySelector('.again'),
  document.querySelector('.check'),
  document.querySelector('.highscore')
);

/* Help with debugging */
console.log(game.hiddenNumber);
