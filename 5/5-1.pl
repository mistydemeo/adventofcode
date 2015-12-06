{
    local $/=undef;
    open FILE, "input.txt";
    $input = <FILE>;
    close FILE;
}

my $nice = 0;

foreach my $line (split("\n", $input)) {
    if ($line =~ /(ab|cd|pq|xy)/) {
        next;
    }
    my $count = ($line =~ tr/[aeiou]//);
    if ($count < 3) {
        next;
    }
    if ($line !~ /(.)\1/) {
        next;
    }
    $nice++;
}

print $nice."\n";
