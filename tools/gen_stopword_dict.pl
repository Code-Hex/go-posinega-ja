#!/usr/bin/env perl

use strict;
use warnings;
use utf8;
use feature qw/say/;

my $s = "var stopWords = map[string]struct{}{\n";

# wget http://svn.sourceforge.jp/svnroot/slothlib/CSharp/Version1/SlothLib/NLP/Filter/StopWord/word/Japanese.txt
while (<>) {
    $_ =~ s/\x{000d}\x{000a}$|\x{000d}$|\x{000a}$//g;
    $_ =~ s/[\r\n]+\z//;
    my $text = $_;
    $s .= "\"$text\": struct{}{},\n" if $text ne "";
}

$s .= "}\n";

print $s;
