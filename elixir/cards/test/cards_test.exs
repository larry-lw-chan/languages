defmodule CardsTest do
  use ExUnit.Case
  doctest Cards

  # test "Create a deck of cards" do
  #   assert Cards.create_deck() == ["Ace of Spades", "Two of Spades", "Three of Spades", "Four of Spades",
  #   "Five of Spades", "Ace of Hearts", "Two of Hearts", "Three of Hearts",
  #   "Four of Hearts", "Five of Hearts", "Ace of Clubs", "Two of Clubs",
  #   "Three of Clubs", "Four of Clubs", "Five of Clubs", "Ace of Diamonds",
  #   "Two of Diamonds", "Three of Diamonds", "Four of Diamonds", "Five of Diamonds"]
  # end

  test "create_deck makes 20 cards" do
    deck_length = length(Cards.create_deck())
    assert(deck_length == 20)
  end

  test "Shuffling a deck randomizes it" do
    deck = Cards.create_deck()
    shuffled_deck = Cards.shuffle(deck)
    refute(deck == shuffled_deck)
  end
end
