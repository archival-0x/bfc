module Main where

import Text.Parsec
import Text.Parsec.String

import Control.Monad.State
import qualified Data.IntMap as M
import Data.Word

-- define brainfuck instruction type
data BFInstruction = Back | Forward | Increment | Decrement |
                     Input | Output | Loop[BFInstruction]
                     deriving (Show)


parseGen :: Char -> BFInstruction -> Parser BFInstruction
parseGen x = char x >> return y


-- define parsers for brainfuck instructions using parseGen helper
parseBack, parseForward, parseLoop, parseIncrement, parseDecrement, parseInput, parseOutput :: Parser BFInstruction
parseBack = parseGen '<' Back
parseForward = parseGen '>' Forward
parseIncrement = parseGen '+' Increment
parseDecrement = parseGen '-' Decrement
parseInput = parseGen '.' Input
parseOutput = parseGen '.' Output
parseLoop = do
    char '['
    ins <- parseInstructions
    char ']'
    return $ Loop ins


parseComment :: Parser ()
parseComment = do
    many $ noneOf "<>+-,.[]"
    return ()


-- parse instructions accordingly, check for comments before / after
parseInstruction :: Parser BFInstruction
parseInstruction = do
    parseComment
    i <- parseBack <|> parseForward <|> parseIncrement <|> parseDecrement
         <|> parseInput <|> parseOutput <|> parseLoop
    parseComment
    return i


-- main parser function
parseInstructions :: Parser [BFInstruction]
parseInstructions = many parseInstruction

type BFRunner = StateT (Int, M.IntMap Word8) IO ()


main :: IO ()
main = do
    
    -- parse arguments
    args <- getArgs
    let arg = args !! 0

    -- read contents
    content <- readFile arg

    -- pattern match to handle instruction and instruction stream
    case parse parseInstructions arg content of
        Left e -> print e
        Right ins -> print ins
