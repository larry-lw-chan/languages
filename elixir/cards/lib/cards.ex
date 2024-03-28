defmodule Cards do
  @moduledoc """
  A module for creating and manipulating decks of cards.
  """

  @doc """
  Returns a list of strings representing a deck of playing cards

  ## Examples

        iex> Cards.create_deck()
        ["Ace of Spades", "Two of Spades", "Three of Spades", "Four of Spades",
        "Five of Spades", "Ace of Hearts", "Two of Hearts", "Three of Hearts",
        "Four of Hearts", "Five of Hearts", "Ace of Clubs", "Two of Clubs",
        "Three of Clubs", "Four of Clubs", "Five of Clubs", "Ace of Diamonds",
        "Two of Diamonds", "Three of Diamonds", "Four of Diamonds", "Five of Diamonds"]

  """
  def create_deck do
    values = ["Ace", "Two", "Three", "Four", "Five"]
    suits = ["Spades", "Hearts", "Clubs", "Diamonds"]

    for suit <- suits, value <- values, do: "#{value} of #{suit}"
  end

  @doc """
  Shuffles the deck of cards
  """
  def shuffle(deck) do
    Enum.shuffle(deck)
  end


  @doc """
  Returns true if the deck contains the hand, false otherwise

  ## Examples

        iex> deck = Cards.create_deck
        iex> Cards.contains?(deck, "Ace of Spades")
        true

  """
  def contains?(deck, hand) do
    Enum.member?(deck, hand)
  end


  @doc """
  Divide a deck into a hand and the reminader of the deck.
  The 'hand_size' argument determines how many cards should
  be in the hand

  ## Examples

        iex> deck = Cards.create_deck
        iex> {hand, _deck} = Cards.deal(deck, 1)
        iex> hand
        ["Ace of Spades"]

  """
  def deal(deck, hand_size) do
    Enum.split(deck, hand_size)
  end

  @doc """
  Saves a deck to a file
  """
  def save(deck, filename) do
    binary = :erlang.term_to_binary(deck)
    File.write(filename, binary)
  end

  @doc """
  Loads a deck from a file
  """
  def load(filename) do
    # {status, binary} = File.read(filename)
    case File.read(filename) do
      { :ok, binary } -> :erlang.binary_to_term(binary)
      { :error, _reason } -> "That file does not exist"
    end
  end

  @doc """
  A helper function for creating a dealt hand of randomized cards
  """
  def create_hand(hand_size) do
    create_deck() |> shuffle() |> deal(hand_size)
  end
end
