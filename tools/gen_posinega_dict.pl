#!/usr/bin/env perl

use strict;
use warnings;
use utf8;
use feature qw/say/;


my $str = "var posinega = map[string]float64{\n";

my $store = +{};

while (<>) {
    $_ =~ s/\x{000d}\x{000a}$|\x{000d}$|\x{000a}$//g;
    $_ =~ s/[\r\n]+\z//;
    my $text = $_;
    my ($key1, $key2, $not_use, $score) = split /:/, $text;
    $str .= "\"$key1\": $score,\n" unless $store->{$key1};
    $store->{$key1} = 1;
    $str .= "\"$key2\": $score,\n" unless $store->{$key2};
    $store->{$key2} = 1;
}

$str .= "}\n";

print $str;

